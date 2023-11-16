"use client";

import styles from "./page.module.css";
import Button from '@mui/material/Button';
import { useRouter } from "next/navigation";
import SignIn from "../app/components/SignIn";
export default function Home() {
  const router = useRouter();

  return (
    <main className={styles.main}>
      <div className={styles.description}>
        <SignIn />
        <Button variant="contained" onClick={() => router.push("/test1")}>TEST1</Button>

      </div>
    </main>
  );
}
