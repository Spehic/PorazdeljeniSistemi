#!/bin/bash
#SBATCH --nodes=1
#SBATCH --array=0-4
#SBATCH --reservation=fri
#SBATCH --output=process-%a.txt

module load Go
go build .
srun go run TretjaDomaca.go -p 9200 -id $SLURM_ARRAY_TASK_ID -n $SLURM_ARRAY_TASK_COUNT -m 5 -k 2
~/go/bin/GoVector --log_type shiviz --log_dir . --outfile Log-Full.log
