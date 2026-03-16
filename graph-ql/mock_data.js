// MongoDB mock data for graphql-job-board database, "jobs" collection.
// Run this in DataGrip against the graphql-job-board database.

db.jobs.insertMany([
    {
        title: "Senior Backend Engineer",
        description: "Design and build scalable microservices using Go, gRPC, and Kubernetes. Experience with distributed systems required.",
        company: "TechNova",
        url: "https://technova.io/careers/senior-backend"
    },
    {
        title: "Full Stack Developer",
        description: "Build modern web applications with React, TypeScript, and GraphQL on the frontend, and Node.js on the backend.",
        company: "WebCraft",
        url: "https://webcraft.dev/jobs/fullstack"
    },
    {
        title: "DevOps Engineer",
        description: "Manage CI/CD pipelines, container orchestration with Docker and Kubernetes, and cloud infrastructure on AWS.",
        company: "CloudBridge",
        url: "https://cloudbridge.com/careers/devops"
    },
    {
        title: "Machine Learning Engineer",
        description: "Develop and deploy ML models for recommendation systems. Proficiency in Python, TensorFlow, and MLOps practices.",
        company: "DataMinds",
        url: "https://dataminds.ai/jobs/ml-engineer"
    },
    {
        title: "Frontend Developer",
        description: "Create pixel-perfect, accessible UIs using Vue.js and Tailwind CSS. Strong understanding of web performance optimization.",
        company: "PixelForge",
        url: "https://pixelforge.design/careers/frontend"
    },
    {
        title: "Site Reliability Engineer",
        description: "Ensure 99.99% uptime for production services. Experience with monitoring, alerting, and incident response.",
        company: "UpGuard Systems",
        url: "https://upguardsystems.com/jobs/sre"
    },
    {
        title: "Blockchain Developer",
        description: "Build smart contracts and decentralized applications on Ethereum and Solana. Solidity and Rust experience preferred.",
        company: "ChainLabs",
        url: "https://chainlabs.xyz/careers/blockchain-dev"
    },
    {
        title: "Mobile Developer",
        description: "Develop cross-platform mobile apps using Flutter and Dart. Experience with native iOS/Android development is a plus.",
        company: "AppVenture",
        url: "https://appventure.io/jobs/mobile-dev"
    },
    {
        title: "Data Engineer",
        description: "Build and maintain ETL pipelines using Apache Spark, Airflow, and BigQuery for large-scale data processing.",
        company: "DataMinds",
        url: "https://dataminds.ai/jobs/data-engineer"
    },
    {
        title: "Security Engineer",
        description: "Conduct security audits, penetration testing, and implement security best practices across the engineering organization.",
        company: "CyberShield",
        url: "https://cybershield.sec/careers/security-engineer"
    }
]);
