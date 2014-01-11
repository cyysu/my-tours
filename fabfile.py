from fabric.api import *

# the user to use for the remote commands
env.user = 'minimum'
# the servers where the commands are executed
env.hosts = ['192.241.196.189']

def deploy():
    with cd("go/src/code.google.com/p/go-tour"):
        run('git pull')
        run('pkill gotour', warn_only=True)
        run('./bin/gotour', pty=False)
