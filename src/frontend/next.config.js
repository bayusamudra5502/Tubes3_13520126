/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  env: {
    BACKEND_sERVER: process.env.BACKEND_sERVER,
  }
}

module.exports = nextConfig
