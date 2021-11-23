

import json
import pickle


def fun(name,password):
    #deserialization untrusted data is always dangerous
    #since this functionality can be done with simple json file
    #thus switiching to it will the application secure.

    s = {"username":name,"password":password}

    with open('users.json',"wt") as fp:
        json.dump(s,fp)




if __name__ == '__main__':
    #u = input("Username : ")
    #p = input("Password : ")
    u = "duckie"
    p = "password"
    yo_fun = fun(u,p)
