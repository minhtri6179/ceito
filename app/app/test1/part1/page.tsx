import Head from "next/head";
import Image from "next/image";
import ImageGallery from "./ImageGallery";
import Navbar from "../navbar/navbar";
export default function Test1() {
  return (
    <>
      <Navbar />
      <ImageGallery />
    </>
  );
}
