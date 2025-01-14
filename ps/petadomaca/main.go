package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"PetaDomaca/myCuda"
	"os"
	"unsafe"
	"time"

	"github.com/InternatBlackhole/cudago/cuda"
)

func rgbaToGray(img image.Image) *image.Gray {
	var (
		bounds = img.Bounds()
		gray   = image.NewGray(bounds)
	)
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			gray.Set(x, y, color.GrayModel.Convert(img.At(x, y)))
		}
	}
	return gray
}

func cpuProcess(image *image.Gray , width, height int ) {
	
	var medArr [9]uint8
	img_in := make([]uint8, len(image.Pix))
	copy(img_in, image.Pix)

	fmt.Println(width, height)
	for ipx := 0 ; ipx < width * height; ipx++ {
		//i am on the left wall
        if ipx % width == 0  {
            // if i am on top left
            if  ipx - width < 0  { 
				medArr[0] = img_in[ipx] 
			} else { 
				medArr[0] = img_in[ ipx - width ]
			}

            medArr[3] = img_in[ipx]

            //if a am on bottom left
            if  ipx + width >= width * height  { 
				medArr[6] = img_in[ipx]
			} else {
				medArr[6] = img_in[ ipx + width ]
			}
		// i am not on left wall
        } else {
			//if i am on top
			if ipx < width {
				medArr[0] = img_in[ ipx - 1 ]
			} else { 
				medArr[0] = img_in[ ipx - width - 1 ]
			}
            
            medArr[3] = img_in[ ipx - 1 ]

			//if i am on bottom
			if ipx + width >= width * height {
				medArr[6] = img_in[ ipx - 1]
			} else {
				medArr[6] = img_in[ ipx + width - 1]
			}
            
        }

        //i am on the right wall
        if ipx % width == width - 1  {
            // if i am on top right
            if  ipx - width < 0 { 
				medArr[2] = img_in[ipx]
			} else { 
				medArr[2] = img_in[ ipx - width ]; 
			}

            medArr[5] = img_in[ipx];

            //if a am on bottom right
            if  ipx + width > width * height {  
				medArr[8] = img_in[ipx] 
			} else { 
				medArr[8] = img_in[ ipx + width ] 
			}
		//i am not on the right wall
        } else {
			//if i am on top
			if ipx < width {
				medArr[2] = img_in[ ipx + 1 ]
			} else { 
				medArr[2] = img_in[ ipx - width + 1 ]
			}
            
            medArr[5] = img_in[ ipx + 1 ]

			//if i am on bottom
			if ipx + width >= width * height {
				medArr[8] = img_in[ ipx + 1]
			} else {
				medArr[8] = img_in[ ipx + width + 1]
			}
		}

        // i am on top
        if ipx < width  {
            // if i am on top left
            if  ipx - 1 < 0 { 
				medArr[0] = img_in[ipx]; 
			} else { 
				medArr[0] = img_in[ ipx - 1 ] 
			}

            medArr[1] = img_in[ipx];

            //if a am on top right
            if  (ipx + 1) % width == 0  { 
				medArr[2] = img_in[ipx] 
			} else { 
				medArr[2] = img_in[ ipx + 1 ] 
			}
		//i am not on top
        } else {
			//if i am on left wall
			if ipx % width == 0 {
				medArr[0] = img_in[ ipx - width ]
			} else {
				medArr[0] = img_in[ ipx - width - 1 ]
			}
            
            medArr[1] = img_in[ ipx - width ]

			//if i am on right wall
			if  (ipx + 1) % width == 0  {
				medArr[2] = img_in[ ipx - width]	
			} else {
            	medArr[2] = img_in[ ipx - width + 1] 
			}
        }


        // i am on bottom
        if ipx + width >= width * height  {
            // if i am on left wall
            if  ipx % width == 0 { 
				medArr[6] = img_in[ipx] 
			} else { 
				medArr[6] = img_in[ ipx - 1 ] 
			}

            medArr[7] = img_in[ipx]

            //if a am on right wall
            if ipx + 1 >= width * height { 
				medArr[8] = img_in[ipx] 
			} else { 
				medArr[8] = img_in[ ipx + 1 ] 
			}
		//i am not on bottom
        } else {
			//if i am on left wall
			if  ipx % width == 0 {
				medArr[6] = img_in[ ipx + width ]
			} else {
				medArr[6] = img_in[ ipx + width - 1 ]
			}
            
            medArr[7] = img_in[ ipx + width ]

			//if i am on right wall
			if  (ipx + 1) % width == 0  { 
				medArr[8] = img_in[ ipx + width ]
			} else {
				medArr[8] = img_in[ ipx + width + 1 ]
			}
            
        }

        medArr[4] = img_in[ ipx ]
        
        for a := 0; a < 8; a++ {
            for b := 0; b < 9 - a - 1; b++ {
                if (medArr[b] > medArr[b + 1]) {
                    temp := medArr[b];
                    medArr[b] = medArr[b + 1];
                    medArr[b + 1] = temp;
                }
            }
        }


        image.Pix[ipx] = medArr[4];

	}
}

func main() {
	// read command line arguments
	inputImageStr := flag.String("i", "", "input image")
	outputImageStr := flag.String("o", "", "output image")
	flag.Parse()
	if *inputImageStr == "" || *outputImageStr == "" {
		panic("Missing input or output image arguments\nUsage: go run main.go -i input.png -o output.png")
	}

	//Initialize CUDA API on OS thread
	var err error
	dev, err := cuda.Init(0)
	if err != nil {
		panic(err)
	}
	defer dev.Close()

	//Open image file
	inputFile, err := os.Open(*inputImageStr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Read image " + *inputImageStr)
	defer inputFile.Close()

	//Decode image
	inputImage, err := png.Decode(inputFile)
	if err != nil {
		panic(err)
	}
	
	//Convert image to grayscale
	inputImageGray := rgbaToGray(inputImage)

	imgSize := inputImageGray.Bounds().Size()
	size := uint64(imgSize.X * imgSize.Y)

	//Allocate memory on the device for input and output image
	imgInDevice, err := cuda.DeviceMemAlloc(size)
	if err != nil {
		panic(err)
	}
	defer imgInDevice.Free()

	imgOutDevice, err := cuda.DeviceMemAlloc(size)
	if err != nil {
		panic(err)
	}
	defer imgOutDevice.Free()


	deviceStart := time.Now()

	//Copy image to device
	err = imgInDevice.MemcpyToDevice(unsafe.Pointer(&inputImageGray.Pix[0]), size)
	if err != nil {
		panic(err)
	}

	//Specify grid and block size
	dimBlock := cuda.Dim3{X: 16, Y: 16, Z: 1}
	dimGrid := cuda.Dim3{
		X: uint32(imgSize.X/16 + 1),
		Y: uint32(imgSize.Y/16 + 1),
		Z: 1,
	}

	//Call the kernel to execute on the device
	err = myCuda.Process(dimGrid, dimBlock, imgInDevice.Ptr, imgOutDevice.Ptr, int32(imgSize.X), int32(imgSize.Y))
	if err != nil {
		panic(err)
	}

	//Copy image back to host
	imgOutHost := make([]byte, size)
	imgOutDevice.MemcpyFromDevice(unsafe.Pointer(&imgOutHost[0]), size)

	//TIME GPU END
	deviceTime := time.Since(deviceStart)

	//TIME CPU START
	cpuStart := time.Now()
	cpuProcess(inputImageGray, int(imgSize.X), int(imgSize.Y))
	//TIME CPU END
	cpuTime := time.Since(cpuStart)


	//Save gpu image to file
	outputFile, err := os.Create(*outputImageStr)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	outputImage := image.NewGray(inputImageGray.Bounds().Bounds())
	outputImage.Pix = imgOutHost

	err = png.Encode(outputFile, outputImage)
	if err != nil {
		panic(err)
	}

	//Save cpu image to file
	outputFileCpu, err := os.Create("cpu.png")
	if err != nil {
		panic(err)
	}
	defer outputFileCpu.Close()

	err = png.Encode(outputFileCpu, inputImageGray)
	if err != nil {
		panic(err)
	}

	fmt.Println("Image saved to " + *outputImageStr, "and", *outputImageStr + "cpu" )
	fmt.Println( "Cpu time:", cpuTime, " Device time:", deviceTime, " Speedup:", float64(cpuTime)/float64(deviceTime) )
}