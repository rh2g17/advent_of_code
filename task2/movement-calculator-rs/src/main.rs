use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() {
    let answer1 = part1();
    let answer2 = part2();
    println!("Answer 1: {}", answer1);
    println!("Answer 2: {}", answer2);
}

enum Moves {
    Up(i16),
    Down(i16),
    Forward(i16),
}


fn create_move(direction: &str, amount: i16) -> Moves {
    return match direction {
        "forward" => Moves::Forward(amount),
        "up" => Moves::Up(amount),
        "down" => Moves::Down(amount),
        _ => panic!("Not a move!")
    }
}

fn generate_move_list() -> Vec<Moves> {
    let file = File::open("../input").expect("Error opening file");
    let reader = BufReader::new(file);
    let mut moves: Vec<Moves> = Vec::new();

    for line in reader.lines() {
        let line = line.unwrap();
        let mut split = line.split(" ");

        let command = create_move(
            split.next().unwrap(), 
            String::from(split.next().unwrap()).parse::<i16>().unwrap()
        );

        moves.push(command)
    }
    moves
}


fn part1() -> i64 {

    let mut horizontal: i64 = 0;
    let mut depth: i64 = 0;

    let moves = generate_move_list();

    for command in moves {
        match command {
            Moves::Forward(amount) => horizontal += i64::from(amount),
            Moves::Down(amount) => depth += i64::from(amount),
            Moves::Up(amount) => depth -= i64::from(amount),
        }
    }

    return horizontal * depth
}

fn part2() -> i64 {
    
    let mut horizontal: i64 = 0;
    let mut depth: i64 = 0;
    let mut aim: i64 = 0;

    let moves = generate_move_list();

    for command in moves {
        match command {
            Moves::Forward(amount) => {
                horizontal += i64::from(amount);
                depth += i64::from(amount) * aim;
            },
            Moves::Down(amount) => aim += i64::from(amount),
            Moves::Up(amount) => aim -= i64::from(amount),
        }
    }

    return horizontal * depth
}