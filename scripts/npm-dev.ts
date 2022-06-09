const npmDevCommand = ["cmd", "/c", "cd", "./frontend", "&&", "npm", "run", "dev"];

// create subprocess
const processFontendDev = Deno.run({ cmd: npmDevCommand });

// await its completion
await processFontendDev.status();