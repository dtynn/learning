use std::io::{self, BufRead, Write};
use std::process::exit;

fn main() {
    let stdin = io::stdin();
    let stdout = io::stdout();
    let mut stdin_lock = stdin.lock();
    let mut stdout_lock = stdout.lock();
    let mut line_buf = String::new();

    loop {
        print_prompt(&mut stdout_lock);

        stdin_lock.read_line(&mut line_buf).unwrap();

        match line_buf.trim_end() {
            ".exit" => {
                println!("bye!");
                exit(0)
            }
            cmd => println!("Unrecognized comman '{}'", cmd),
        }

        line_buf.clear();
    }
}

fn print_prompt<W: Write>(w: &mut W) {
    w.write_all("db > ".as_bytes())
        .and_then(|_| w.flush())
        .unwrap();
}
