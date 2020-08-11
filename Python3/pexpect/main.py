import sys
from pexpect import pxssh


# child = pexpect.spawn('ssh sadm@care-i.comp.nus.edu.sg')
# child.expect('[\s\S]+ password: ')
# print(child.before)
# print(child.after)
#
# child.sendline('careadmin')
# child.expect('Last login: [\s\S]+')
# print(child.before)
# print(child.after)
#
# child.sendline('cd Documents')
# child.expect('\r\n[\s\S]+')
# print(child.before)
# print(child.after)
#
# child.sendline("ls")
# child.expect('\r\n[\s\S]+')
# print(child.before)
# print(child.after)

def login2():
    try:
        s = pxssh.pxssh()
        s.login('care-i.comp.nus.edu.sg', 'sadm', 'careadmin')
        s.sendline('uptime')  # run a command
        s.prompt()  # match the prompt
        print(s.before.decode("utf-8"))  # print everything before the prompt.
        s.sendline('pwd')
        s.prompt()
        print(s.before.decode('utf-8'))
    except pxssh.ExceptionPxssh as e:
        print("pxssh failed on login.")
        print(e)
    else:
        s.logout()


def login(host: str, username: str, password: str) -> pxssh:
    s = pxssh.pxssh()
    s.login(host, username, password)
    return s


def run_command(cmd: str, sessions: list) -> list:
    result = []
    for s in sessions:
        s.sendline(cmd)
        s.prompt()
        result.append(s.before)
    return result


class Sessions:
    def __init__(self):
        self.__sessions = []

    def add_session(self, host: str, username: str, password: str):
        s = pxssh.pxssh()
        s.login(host, username, password)
        self.__sessions.append(s)

    def run_command(self, cmd) -> list:
        results = []
        for s in self.__sessions:
            s.sendline(cmd)
            s.prompt()
            results.append(s.before)
        return results


if __name__ == '__main__':
    # print(sys.argv)
    # sessions = []
    # sessions.append(login('care-i.comp.nus.edu.sg', 'sadm', 'careadmin'))
    # r = run_command('uptime', sessions)
    # print(r)

    sessions = Sessions()
    sessions.add_session('care-i.comp.nus.edu.sg', 'sadm', 'careadmin')
    r = sessions.run_command('uptime')
    print(r)


