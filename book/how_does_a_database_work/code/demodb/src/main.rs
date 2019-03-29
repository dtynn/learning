use std::io::{self, BufRead, Write};
use std::process::exit;

fn main() {
    let stdin = io::stdin();
    let stdout = io::stdout();
    let mut stdin_lock = stdin.lock();
    let mut ctx = Context {
        stdout: stdout.lock(),
    };
    let mut line_buf = String::new();

    loop {
        ctx.print_prompt().unwrap();

        stdin_lock.read_line(&mut line_buf).unwrap();

        let s = line_buf.trim();
        if !s.is_empty() {
            let cmd = Command::from_str(s);
            cmd.exec(&mut ctx);
        }

        line_buf.clear();
    }
}

enum MetaCommand {
    Exit,
    Unrecognized(String),
}

impl MetaCommand {
    fn from_str(cmd: &str) -> Self {
        match cmd.to_lowercase().as_str() {
            ".exit" => MetaCommand::Exit,
            _ => MetaCommand::Unrecognized(cmd.to_owned()),
        }
    }

    fn exec(&self, ctx: &mut Context) {
        match self {
            MetaCommand::Exit => {
                ctx.write_line_to_out("bye!").unwrap();
                exit(0);
            }

            MetaCommand::Unrecognized(raw) => {
                ctx.write_line_to_out(&format!("Unrecognized meta command '{}'", raw))
                    .unwrap();
            }
        }
    }
}

enum Statement {
    Select,
    Insert,
    Unrecognized,
}

impl Statement {
    fn from_str(s: &str) -> Self {
        let st = s.to_lowercase();
        if st.starts_with("select") {
            return Statement::Select;
        }

        if st.starts_with("insert") {
            return Statement::Insert;
        }

        Statement::Unrecognized
    }

    fn exec(&self, ctx: &mut Context) {
        match self {
            Statement::Select => {
                ctx.write_line_to_out("Executing SELECT statement").unwrap();
            }

            Statement::Insert => {
                ctx.write_line_to_out("Executing INSERT statement").unwrap();
            }

            Statement::Unrecognized => {
                ctx.write_line_to_out("Unrecognized statement").unwrap();
            }
        }
    }
}

enum Command {
    MetaCommand(MetaCommand),
    Statement(Statement),
}

impl Command {
    fn from_str(s: &str) -> Self {
        if s.starts_with(".") {
            Command::MetaCommand(MetaCommand::from_str(s))
        } else {
            Command::Statement(Statement::from_str(s))
        }
    }

    fn exec(&self, ctx: &mut Context) {
        match self {
            Command::MetaCommand(mc) => mc.exec(ctx),
            Command::Statement(st) => st.exec(ctx),
        }
    }
}

struct Context<'a> {
    stdout: io::StdoutLock<'a>,
}

const CR: [u8; 1] = [b'\n'];
const PROMPT: &str = "db > ";

impl<'a> Context<'a> {
    fn print_prompt(&mut self) -> io::Result<()> {
        self.stdout
            .write_all(PROMPT.as_bytes())
            .and_then(|_| self.stdout.flush())
    }

    fn write_line_to_out(&mut self, s: &str) -> io::Result<()> {
        self.stdout
            .write_all(s.as_bytes())
            .and_then(|_| self.stdout.write_all(&CR[..]))
            .and_then(|_| self.stdout.flush())
    }
}
