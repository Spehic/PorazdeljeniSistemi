#ifdef __cplusplus
extern "C" {
#endif

__global__ void hello(char *message) {
    printf("%s\n", message);
}

__global__ void process(unsigned char *img_in, unsigned char *img_out, int width, int height) {
    // row
    int j = threadIdx.x + blockIdx.x * blockDim.x;
    // col
    int i = threadIdx.y + blockIdx.y * blockDim.y;
    int ipx = i * width + j;

    unsigned char medArr[9];
    while (ipx < width * height) {
        
        //i am on the left wall
        if( ipx % width == 0 ) {
            // if i am on top left
            if ( ipx - width < 0 ) medArr[0] = img_in[ipx];
            else medArr[0] = img_in[ ipx - width ];

            medArr[3] = img_in[ipx];

            //if a am on bottom left
            if ( ipx + width >= width * height ) medArr[6] = img_in[ipx];
            else medArr[6] = img_in[ ipx + width ];
        }else {
			//if i am on top
			if ( ipx < width ) 
				medArr[0] = img_in[ ipx - 1 ];
			 else 
				medArr[0] = img_in[ ipx - width - 1 ];
			
            
            medArr[3] = img_in[ ipx - 1 ];

			//if i am on bottom
			if ( ipx + width >= width * height ) {
				medArr[6] = img_in[ ipx - 1];
			} else {
				medArr[6] = img_in[ ipx + width - 1];
			}
        }

         //i am on the right wall
        if( ipx % width == width - 1 ) {
            // if i am on top right
            if ( ipx - width < 0 ) medArr[2] = img_in[ipx];
            else medArr[2] = img_in[ ipx - width ];

            medArr[5] = img_in[ipx];

            //if a am on bottom right
            if ( ipx + width > width * height ) medArr[8] = img_in[ipx];
            else medArr[8] = img_in[ ipx + width ];
        }else {
			//if i am on top
			if (ipx < width) {
				medArr[2] = img_in[ ipx + 1 ];
			} else { 
				medArr[2] = img_in[ ipx - width + 1 ];
			}
            
            medArr[5] = img_in[ ipx + 1 ];

			//if i am on bottom
			if (ipx + width >= width * height) {
				medArr[8] = img_in[ ipx + 1];
			} else {
				medArr[8] = img_in[ ipx + width + 1];
			}
        }

        // i am on top
        if( ipx < width ) {
            // if i am on top left
            if ( ipx - 1 < 0 ) medArr[0] = img_in[ipx];
            else medArr[0] = img_in[ ipx - 1 ];

            medArr[1] = img_in[ipx];

            //if a am on top right
            if ( (ipx + 1) % width == 0 ) medArr[2] = img_in[ipx];
            else medArr[2] = img_in[ ipx + 1 ];
        }else {
			//if i am on left wall
			if (ipx % width == 0) {
				medArr[0] = img_in[ ipx - width ];
			} else {
				medArr[0] = img_in[ ipx - width - 1 ];
			}
            
            medArr[1] = img_in[ ipx - width ];

			//if i am on right wall
			if  ( (ipx + 1) % width == 0 )  {
				medArr[2] = img_in[ ipx - width ];
			} else {
            	medArr[2] = img_in[ ipx - width + 1];
			}
        }


        // i am on bottom
        if( ipx + width >= width * height ) {
            // if i am on bottom left
            if ( ipx % width == 0 ) medArr[6] = img_in[ipx];
            else medArr[6] = img_in[ ipx - 1 ];

            medArr[7] = img_in[ipx];

            //if a am on bottom right
            if ( ipx + 1 >= width * height ) medArr[8] = img_in[ipx];
            else medArr[8] = img_in[ ipx + 1 ];
        }else {
			//if i am on left wall
			if (ipx % width == 0) {
				medArr[6] = img_in[ ipx + width ];
			} else {
				medArr[6] = img_in[ ipx + width - 1 ];
			}
            
            medArr[7] = img_in[ ipx + width ];

			//if i am on right wall
			if  ((ipx + 1) % width == 0 ) { 
				medArr[8] = img_in[ ipx + width ];
			} else {
				medArr[8] = img_in[ ipx + width + 1 ];
			}
        }

        medArr[4] = img_in[ ipx ];
        
        for (int a = 0; a < 8; a++) {
            for (int b = 0; b < 9 - a - 1; b++) {
                if (medArr[b] > medArr[b + 1]) {
                    int temp = medArr[b];
                    medArr[b] = medArr[b + 1];
                    medArr[b + 1] = temp;
                }
            }
        }


        img_out[ipx] = medArr[4];
        ipx += blockDim.x * gridDim.x * blockDim.y * gridDim.y;
    }
}



#ifdef __cplusplus
}
#endif