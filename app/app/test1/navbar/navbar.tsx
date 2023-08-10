import * as React from "react";
import Stack from "@mui/material/Stack";
import NavLinkButton from "./NavLinkButton";
import AudioPlayer from "./AudioPlayer";
export default function Navbar() {
  
  return (
    <Stack spacing={2}>
      <AudioPlayer audioSrc="/test1/audio.mp3" />
      <Stack direction="row" spacing={3}>
        <NavLinkButton href="/test1/part1" label="Part1" />
        <NavLinkButton href="/test1/part2" label="Part2" />
        <NavLinkButton href="/test1/part3" label="Part3" />
        <NavLinkButton href="/test1/part4" label="Part4" />
        <NavLinkButton href="/test1/part5" label="Part5" />
        <NavLinkButton href="/test1/part6" label="Part6" />
        <NavLinkButton href="/test1/part7" label="Part7" />
        {/* Add more NavLinkButton components for other links */}
      </Stack>
    </Stack>
  );
}
