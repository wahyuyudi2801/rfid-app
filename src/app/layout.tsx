import type { Metadata } from "next";
import { Nunito } from "next/font/google";
import "@/styles/global.css"

const fontSans = Nunito({
  variable: "--font-sans",
  weight: "600",
  subsets: ["latin"],
});

const fontMono = Nunito({
  variable: "--font-mono",
  weight: "600",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: "RFID App",
  description: "Aplikasi absensi RFID",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body
        className={`${fontSans.variable} ${fontMono.variable} antialiased`}
      >
        {children}
      </body>
    </html>
  );
}
