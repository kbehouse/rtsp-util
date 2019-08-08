import os
from os import system
from shutil import copyfile, copy2, copytree

ip_head = '192.168.0.'
ip_tail = 140
mp4_dir = 'videos/'
docker_network_name = 'macvlan-192-168-0'
restart_always = True
restart_always_str = '--restart always' if  restart_always else '--restart no'

files = [f for f in os.listdir(mp4_dir)]
print(files)
for i, f in enumerate(files):
    if  f[-4:]!='.mp4':
        continue

    ip = ip_head + str(ip_tail)
    
    host_video = os.path.join(os.getcwd(), mp4_dir, f)
    docker_name = ip + '_' + f
    print('Broadcast:', host_video)
    cmd = 'docker run -d --network {} --ip={} -v {}:/input_video --name {} {}  kbehouse/video2rtsp'. \
        format(docker_network_name, ip, host_video, docker_name, restart_always_str)

    ip_tail += 1

    print(cmd)
    system(cmd)