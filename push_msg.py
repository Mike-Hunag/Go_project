import os
import subprocess
import json

def push_msg(Epoch, Loss, Accuracy, CER):
    go_post_directory = "C:/Users/mike/Desktop/Python_experiment/CRNN/vlnet_1D/GoPost"
    os.chdir(go_post_directory)
    data = {
        "Epoch": Epoch,
        "Loss": Loss,
        "Accuracy": Accuracy,
        "CER": CER,
    }

    data_json = json.dumps(data)

    go_program = "C:/Users/mike/Desktop/Python_experiment/CRNN/vlnet_1D/GoPost/main.go"
    command = ["go", "run", go_program]

    if os.path.isfile(go_program):
        process = subprocess.Popen(command, stdin=subprocess.PIPE, stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True, shell=True)
    else:
        print("The specified file does not exist:", go_program)
    
    stdout, stderr = process.communicate(input=data_json)

    if process.returncode == 0:
        print(f"Output: {stdout}")
    else:
        print(f"Error: {stderr}")

if __name__ == '__main__':
    push_msg(1, 90.001, 0.0, 0.95)