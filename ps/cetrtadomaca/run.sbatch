#!/bin/bash
#SBATCH --nodes=1
#SBATCH --array=0-1
#SBATCH --reservation=fri
#SBATCH --output=process-%a.txt

module load Go
go build .
srun go run Cetrtadomaca.go -s localhost -p 9205 -id $SLURM_ARRAY_TASK_ID -n $SLURM_ARRAY_TASK_COUNT