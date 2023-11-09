const fs = require("fs");
const path = require("path");

function splitImageName(imageName) {
  let res = imageName.split(".")[0].split("-");
  return res;
}
function getImage(test_part) {
  const folderPath = test_part;

  // Read all files in the folder
  const files = fs.readdirSync(folderPath);
  let res = {};
  // Log the list of file names
  console.log("File names in the folder:");
  files.forEach((file) => {
    if (file.charAt(0) != ".") {
      res[file] = splitImageName(file);
    }
  });
  return res;
}

let res = getImage("../public/test1/part3");
console.log(res["62-64.png"]);
console.log(res);
