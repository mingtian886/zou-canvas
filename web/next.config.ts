import type { NextConfig } from "next";
import { PHASE_DEVELOPMENT_SERVER } from "next/constants";

export default function nextConfig(phase: string): NextConfig {
    return {
        output: "standalone",
        allowedDevOrigins: phase === PHASE_DEVELOPMENT_SERVER ? ["*.*.*.*"] : [],
        typescript: {
            ignoreBuildErrors: true,
        },
    };
}
