const goRunCommand = ["go", "run", "."];

// create subprocess
const processGoRun = Deno.run({ cmd: goRunCommand });

// await its completion
await processGoRun.status();