function getAllDataFromLocalStorage() {
  const localStorageData = {};
  for (let i = 0; i < localStorage.length; i++) {
    const key = localStorage.key(i);
    const value = localStorage.getItem(key);
    localStorageData[key] = value;
  }
  return localStorageData;
}
function isLetter(c) {
  return c.toLowerCase() != c.toUpperCase();
}
function submit() {
  return new Promise((resolve, reject) => {
    const dataObject = getAllDataFromLocalStorage();
    const questions = [];
    const answers = [];
    const part1 = dataObject["selectedAnswers"].split("");
    let idx = 0;
    for (let i = 0; i < part1.length; i++) {
      if (isLetter(part1[i])) {
        questions.push(idx + 1);
        answers.push(part1[i]);
        idx += 1;
      }
    }

    for (let i = 6; i < 200; i++) {
      const currentidx = "selectedAnswers-" + i;
      questions.push(i + 1);
      answers.push(dataObject[currentidx]);
    }

    const answers_idx = [];
    for (let i = 0; i < answers.length; i++) {
      let curidx;
      if (answers[i] == "A") {
        curidx = i * 4 + 1;
      } else if (answers[i] == "B") {
        curidx = i * 4 + 2;
      } else if (answers[i] == "C") {
        curidx = i * 4 + 3;
      } else if (answers[i] == "D") {
        curidx = i * 4 + 4;
      } else {
        curidx = 0;
      }
      answers_idx.push(curidx);
    }
    fetch("https://ceito.onrender.com/submit", {
      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ question_id: questions, answer_id: answers_idx }),
    })
      .then((response) => response.json())
      .then((parsedResponse) => {
        resolve(parsedResponse["your_score"]); // Resolve the promise with the parsed response
      })
      .catch((error) => {
        reject(error); // Reject the promise with an error if there's an issue
      });
  });
}
export default submit;
