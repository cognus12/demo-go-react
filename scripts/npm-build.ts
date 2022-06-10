const npmBuildCommand = ["cmd", "/c", "cd", "./frontend", "&&", "npm", "run", "build"];

// create subprocess
const processFontendBuild = Deno.run({ cmd: npmBuildCommand });

// await its completion
await processFontendBuild.status();