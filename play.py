import keyboard
import time
import sys
import configparser
import socket

def is_port_in_use(port):
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
        return s.connect_ex(('localhost', port)) == 0


def init_server(server, environment, build):
    # open a new terminal
    keyboard.press_and_release('ctrl + shift + `')
    time.sleep(1)

    # change the directory
    to_write = "cd " + server['path']
    keyboard.write(to_write)
    keyboard.press_and_release('enter')


    # Determine the command based on environment and build flag
    if build:
        command_key = 'build_command'
    else:
        command_key = f'{environment}_command'
        if 'depends_on_port' in server:
            while not is_port_in_use(int(server['depends_on_port'])):
                print("Waiting for port " + server['depends_on_port'] + " to be in use.")
                time.sleep(1)

    # run the command
    to_write = server[command_key]
    keyboard.write(to_write)
    keyboard.press_and_release('enter')

    print("Server " + server['name'] + " started.")

def parse_ini(ini_path):
    config = configparser.ConfigParser()
    config.read(ini_path)

    servers = []
    environment = config['DEFAULT'].get('environment', 'dev')
    build = config['DEFAULT'].getboolean('build', fallback=False)
    
    for section in config.sections():
        server_info = {
            'name': section,
            'path': config[section]['path'],
            'dev_command': config[section]['dev_command'],
            'build_command': config[section]['build_command'],
            'prod_command': config[section]['prod_command'],
        }
        # Check for 'depends_on_port' key
        if 'depends_on_port' in config[section]:
            server_info['depends_on_port'] = config[section]['depends_on_port']

        servers.append((server_info, environment, build))

    return servers


def main():
    servers = parse_ini(ini_file_path)
    for server_info, environment, build in servers:
        init_server(server_info, environment, build)



if __name__ == "__main__":
    if len(sys.argv) > 1:
        ini_file_path = sys.argv[1]
        main()
    else:
        print("No INI file provided.")