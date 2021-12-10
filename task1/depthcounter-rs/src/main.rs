use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() {
    let answer1 = part1();
    let answer2 = part2();
    println!("Answer 1: {}", answer1);
    println!("Answer 2: {}", answer2);
}

fn part1() -> u16 {
    let file = File::open("../input").expect("Error opening file");
    let reader = BufReader::new(file);
    let mut measurements = Vec::new();

    for line in reader.lines() {
        let line = line.unwrap();
        measurements.push(line.parse::<i32>().unwrap());
    }

    let mut number_of_increases: u16 = 0;

    let previous = measurements.get(0);
    let mut previous = match previous {
        None => panic!("Issue getting a measurement"),
        Some(value) => *value,
    };

    for measurement in measurements {
        if measurement > previous {
            number_of_increases += 1
        }
        previous = measurement
    }

    return number_of_increases
}

fn part2() -> u16 {
    let file = File::open("../input").expect("Error opening file");
    let reader = BufReader::new(file);
    let mut measurements = Vec::new();

    for line in reader.lines() {
        let line = line.unwrap();
        measurements.push(line.parse::<i32>().unwrap());
    }

    let mut number_of_increases: u16 = 0;

    let mut previous = measurements.get(0).unwrap() + measurements.get(1).unwrap() + measurements.get(2).unwrap();

    for i in 0..(measurements.len() - 2) {
        let current = measurements[i] + measurements[i + 1] + measurements[i + 2];
        if current > previous {
            number_of_increases += 1
        }

        previous = current
    }

    return number_of_increases
}
