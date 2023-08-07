// pages/ImageGallery.js
"use client";
import React, { useState } from "react";
import {
  FormControl,
  FormControlLabel,
  FormLabel,
  RadioGroup,
  Radio,
} from "@mui/material";

const ImageGallery = () => {
  // Replace 'path_to_images_folder' with the actual path where your images are located
  const imagesFolder = "test1/part1";

  // Get a list of all image files in the specified folder
  const imageFiles = [
    // Add your image filenames here or use a loop to fetch from the folder dynamically
    "q1.png",
    "q2.png",
    "q3.png",
    "q4.png",
    "q5.png",
    "q6.png",
    // Add more images as needed...
  ];
  const [selectedImageIndex, setSelectedImageIndex] = useState(0);
  const handleImageChange = (index) => {
    setSelectedImageIndex(index);
  };

  return (
    <div>
      {imageFiles.map((imageFile, index) => (
        <div key={index}>
          <FormControl>
            <FormLabel id="qus-ans-ansersheet">Question {index + 1}</FormLabel>
            <RadioGroup
              aria-labelledby="demo-radio-buttons-group-label"
              name="ansersheet"
            >
              <FormControlLabel value="A" control={<Radio />} label="A" />
              <FormControlLabel value="B" control={<Radio />} label="B" />
              <FormControlLabel value="C" control={<Radio />} label="C" />
              <FormControlLabel value="D" control={<Radio />} label="D" />
            </RadioGroup>
          </FormControl>

          <img
            key={index}
            src={`/${imagesFolder}/${imageFile}`}
            alt={`Image ${index + 1}`}
            style={{ margin: "5px", marginBottom: "20px" }}
            width={500}
            height={500}
          />
        </div>
      ))}
    </div>
  );
};

export default ImageGallery;
