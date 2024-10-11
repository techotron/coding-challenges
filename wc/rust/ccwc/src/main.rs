use clap::Parser;
use std::fs::File;
use std::io::{Result, Read, BufReader, BufRead};
use std::path::Path;

#[derive(Parser, Debug)]
#[command(version, about, long_about = None)]
struct Args {
    #[arg(short, long)]
    count: bool,

    #[arg(short, long)]
    lines: bool,

    #[arg(short, long)]
    words: bool,

    #[arg(value_name = "FILENAME")]
    filename: String,
}

fn main() -> Result<()> {
    let args = Args::parse();
    
    let path = Path::new(&args.filename);
    let mut file = File::open(&args.filename)?;

    if args.count {
        let mut contents = Vec::new();
        file.read_to_end(&mut contents)?;
        println!("{} {}", contents.len(), path.file_name().unwrap().to_str().unwrap());
    }

    if args.lines {
        let reader = BufReader::new(&file);
        let line_count = reader.lines().count();
        println!("{} {}", line_count, path.file_name().unwrap().to_str().unwrap());
    }

    if args.words {
        let reader = BufReader::new(&file);
        let word_count = reader.lines()
            .map(|line| line.unwrap().split_whitespace().count())
            .sum::<usize>();
        println!("{} {}", word_count, path.file_name().unwrap().to_str().unwrap());
    }
    Ok(())
}
