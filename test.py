import subprocess

cases = ['hello', 'HELLO', 'HeLlo HuMaN', '1Hello 2There',
         'Hello\\nThere', '{Hello & There #}', 'hello There 1 to 2!',
         'MaD3IrA&LiSboN']

for case in cases:
    result = subprocess.check_output(['./ascii-art', case], stderr = subprocess.STDOUT, text = True)
    print(result)