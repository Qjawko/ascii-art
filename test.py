import subprocess

cases = ['"hello"', '"HELLO"', '"HeLlo HuMaN"', ]
subprocess.check_output(['./ascii-art', 'hello'])