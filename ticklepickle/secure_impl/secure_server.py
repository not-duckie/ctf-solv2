import json
import pickle

def reverse_fun():
      with open("users.json","rb") as f:
          data = f.read()
          creds = json.loads(data)

      return creds


if __name__ == '__main__':
      print(reverse_fun())
