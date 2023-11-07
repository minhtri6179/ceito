"use client";

import styles from "./page.module.css";
import Button from '@mui/material/Button';
import { useRouter } from "next/navigation";
export default function Home() {
  const router = useRouter();

  return (
    <main className={styles.main}>
      <div className={styles.description}>
        <p>
          Get started by editing&nbsp;
          <code className={styles.code}>app/page.tsx</code>
          
        </p>
        <Button variant="contained" onClick={() => router.push("/test1")}>TEST1</Button>

      </div>
    </main>
  );
}
