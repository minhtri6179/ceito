import requests
import pandas as pd
import numpy as np


class CreateTestDB:
    def __init__(self, url) -> None:
        self.url = url

    def create_question(self, part, test, question_text):
        pass

    def create_answer(self, part, test, question_id):
        pass

    def create_part1(self):
        for i in range(6):
            myobj = {
                "question_text": f"Part1-{i+1}",
                "test_name": "ETS-23-Test1-Part1",
            }
            x = requests.post(self.url, json=myobj)

            print(x.text)

    def create_part2(self):
        for i in range(6, 31):
            myobj = {
                "question_text": f"Part2-{i+1}",
                "test_name": "ETS-23-Test1-Part2",
            }
            x = requests.post(self.url, json=myobj)

            print(x.text)

    def create_part3(self):
        for i in range(31, 69):
            myobj = {
                "question_text": f"Part3-{i+1}",
                "test_name": "ETS-23-Test1-Part3",
            }
            x = requests.post(self.url, json=myobj)

            print(x.text)

    def create_part4(self):
        for i in range(70, 100):
            myobj = {
                "question_text": f"Part4-{i+1}",
                "test_name": "ETS-23-Test1-Part4",
            }
            x = requests.post(self.url, json=myobj)

            print(x.text)

    def create_part5(self):
        for i in range(31, 71):
            myobj = {
                "question_text": f"Part4-{i+1}",
                "test_name": "ETS-23-Test1-Part4",
            }
            x = requests.post(self.url, json=myobj)

            print(x.text)

    def create_part6(self):
        for i in range(31, 71):
            myobj = {
                "question_text": f"Part4-{i+1}",
                "test_name": "ETS-23-Test1-Part4",
            }
            x = requests.post(self.url, json=myobj)

            print(x.text)

    def create_part7():
        for i in range(31, 71):
            myobj = {
                "question_text": f"Part4-{i+1}",
                "test_name": "ETS-23-Test1-Part4",
            }
            x = requests.post(self.url, json=myobj)

            print(x.text)


if __name__ == "__main__":
    df = pd.read_csv("ETS23-Test1.csv")
    question_text = df["Question_text"][::4]
    answer_text = df["Answer_text"]
    no_question = df["Question_id"][::4].astype(int)
