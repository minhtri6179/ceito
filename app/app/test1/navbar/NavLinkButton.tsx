"use client";

import * as React from "react";
import Button from "@mui/material/Button";
import Link from "next/link";
import { usePathname } from 'next/navigation'

interface NavLinkButtonProps {
  href: string;
  label: string;
}

const NavLinkButton: React.FC<NavLinkButtonProps> = ({ href, label }) => {
    const pathname = usePathname()

  return (
    
    <Link href={href} >
      <Button variant={pathname === href ? "contained" : "outlined"}>
        {label}
      </Button>
    </Link>
  );
};

export default NavLinkButton;