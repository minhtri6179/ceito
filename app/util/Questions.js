async function getQuestions() {
  const response = await fetch("http:127.0.0.1:8080/questions", {
    method: "GET",
    mode: "cors",
    headers: {
      "Content-Type": "application/json",
    },
  });
  const questions = await response.json();
  var dictstring = JSON.stringify(questions);
  var fs = require("fs");
  fs.writeFile("questions.json", dictstring, function (err, result) {
    if (err) console.log("error", err);
  });
}
getQuestions();
