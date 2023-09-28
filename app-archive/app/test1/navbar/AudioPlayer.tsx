"use client";
import React, { useState } from "react";

interface AudioPlayerProps {
  audioSrc: string;
}

const AudioPlayer: React.FC<AudioPlayerProps> = ({ audioSrc }) => {
  return (
    <div>
      <audio controls src={audioSrc} style={{ width: "50%", height: "60px" }}/>
    </div>
  );
};

export default AudioPlayer;
