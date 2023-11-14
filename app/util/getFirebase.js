function getFirebase(path) {
  // let q = encodeURIComponent("/test1/part1/q" + i + ".png");

  let baseURL =
    "https://firebasestorage.googleapis.com/v0/b/cieto-20a9b.appspot.com/o/public";
  let img = "?alt=media";
  let q = encodeURIComponent(path);

  return baseURL + q + img;
}

export default getFirebase;
