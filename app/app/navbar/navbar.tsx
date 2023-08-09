import Image from "next/image";
import Styles from "./navbar.module.css";
import Link from "next/link";
export default function Navbar() {
  return (
    <nav className={Styles.nav}>
      <Link href="/">
        <Image
          className={Styles.logo}
          src="/vercel.svg"
          alt="logo"
          width="100"
          height="100"
        />
      </Link>
    </nav>
  );
}
