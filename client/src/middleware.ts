import { NextRequest, NextResponse } from "next/server";

// Add whatever paths you want to PROTECT here
const authRoutes = ["/games/*"];

// Function to match the * wildcard character
function matchesWildcard(path: string, pattern: string): boolean {
  if (pattern.endsWith("/*")) {
    const basePattern = pattern.slice(0, -2);
    return path.startsWith(basePattern);
  }
  return path === pattern;
}

export async function middleware(request: NextRequest) {
  // Shortcut for our login path redirect
  // Note: you must use absolute URLs for middleware redirects
  const LOGIN = `${process.env.NEXT_PUBLIC_BASE_URL}`;

  if (
    authRoutes.some((pattern) =>
      matchesWildcard(request.nextUrl.pathname, pattern)
    )
  ) {
    const user = request.cookies.get("user");
    const userParsed = user ? JSON.parse(user.value) : null;

    // If no token exists, redirect to login
    if (!userParsed) {
      return NextResponse.redirect(LOGIN);
    }
  }

  // Redirect login to app if already logged in
  if (request.nextUrl.pathname === "/") {
    const user = request.cookies.get("user");
    const userParsed = user ? JSON.parse(user.value) : null;

    if (userParsed) {
      return NextResponse.redirect(`${process.env.NEXT_PUBLIC_BASE_URL}/games`);
    }
  }

  return NextResponse.next();
}
