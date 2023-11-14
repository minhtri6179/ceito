import * as React from "react";
import Stack from "@mui/material/Stack";
import AudioPlayer from "./AudioPlayer";
import getFirebase from "../../../util/getFirebase";

export default function Navbar() {
  
  return (
      <AudioPlayer audioSrc={getFirebase("/test1/audio.mp3")} />
  );
}
