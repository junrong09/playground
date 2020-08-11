import sys
from pexpect import pxssh


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
    sessions = Sessions()
    sessions.add_session('care-i.comp.nus.edu.sg', 'sadm', 'careadmin')
    r = sessions.run_command('uptime')
    print(r)


