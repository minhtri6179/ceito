const fs = require("fs");
const { type } = require("os");
const path = require("path");

function splitImageName(imageName) {
  let res = imageName.split(".")[0].split("-");
  return res;
}
function getImage(test_part) {
  try {
    const folderPath = test_part;

    // Read all files in the folder
    const files = fs.readdirSync(folderPath);
    let res = [];

    // Log the list of file names
    console.log("File names in the folder:");

    files.forEach((file) => {
      if (file.charAt(0) !== ".") {
        res.push({
          fileName: file,
          splitParts: splitImageName(file),
        });
      }
    });
    return res;
  } catch (error) {
    console.error("Error reading files:", error);
    return []; // Return an empty array or handle the error accordingly
  }
}

export default getImage;
