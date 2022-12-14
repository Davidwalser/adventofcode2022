use std::str::Lines;

#[derive(Debug, PartialEq)]
struct Rock{
    name: &'static str,
}
#[derive(Debug, PartialEq)]
struct Paper{
    name: &'static str,
}
#[derive(Debug, PartialEq)]
struct Scissor{
    name: &'static str,
}

const ROCK: &str = "Rock";
const PAPER: &str = "Paper";
const SCISSOR: &str = "Scissor";
const WINNING_POINTS: i32 = 6;
const DRAW_POINTS: i32 = 3;
const LOSING_POINTS: i32 = 0;

trait GameObject{
    fn get_points(&self) -> i32;
    fn get_name(&self) -> &str;
    fn wins_against(&self, other: &dyn GameObject) -> bool;
    fn calculate_points(&self, other: &dyn GameObject) -> i32;
    fn convert(&self, input: &str) -> &dyn GameObject;
}

fn detect_winning_points(player1: &dyn GameObject, player2: &dyn GameObject) -> i32{
    if player1.wins_against(player2) {
        return WINNING_POINTS;
    }
    else if player2.wins_against(player1) {
        return LOSING_POINTS;
    }
    return DRAW_POINTS;
}

impl GameObject for Rock{
    fn get_points(&self) -> i32{
        return 1;
    }

    fn get_name(&self) -> &str {
        return self.name;
    }

    fn wins_against(&self, other: &dyn GameObject) -> bool {
        if other.get_name() == SCISSOR {
            return true;
        }
        return false;
    }

    fn calculate_points(&self, other: &dyn GameObject) -> i32 {
        return self.get_points() + detect_winning_points(self,other);
    }

    fn convert(&self, input: &str) -> &dyn GameObject {
        match input {
            "X" => return &Scissor{name: SCISSOR},
            "Y" => return &Rock{name: ROCK},
            "Z" => return &Paper{ name: PAPER},
            _ => panic!("Type not convertable")
        }
    }
}

// X -> Lose
// Y -> Draw
// Z -> Win
impl GameObject for Paper{
    fn get_points(&self) -> i32{
        return 2;
    }

    fn get_name(&self) -> &str {
        return self.name;
    }

    fn wins_against(&self, other: &dyn GameObject) -> bool {
        if other.get_name() == ROCK {
            return true;
        }
        return false;
    }

    fn calculate_points(&self, other: &dyn GameObject) -> i32 {
        return self.get_points() + detect_winning_points(self,other);
    }

    fn convert(&self, input: &str) -> &dyn GameObject {
        match input {
            "X" => return &Rock{name: ROCK},
            "Y" => return &Paper{name: PAPER},
            "Z" => return &Scissor{ name: SCISSOR},
            _ => panic!("Type not convertable")
        }
    }
}

impl GameObject for Scissor{
    fn get_points(&self) -> i32{
        return 3;
    }

    fn get_name(&self) -> &str {
        return self.name;
    }

    fn calculate_points(&self, other: &dyn GameObject) -> i32 {
        return self.get_points() + detect_winning_points(self,other);
    }

    fn wins_against(&self, other: &dyn GameObject) -> bool {
        if other.get_name() == PAPER {
            println!("SCISSOR won against ROCK");
            return true;
        }
        return false;
    }

    fn convert(&self, input: &str) -> &dyn GameObject {
        match input {
            "X" => return &Paper{name: PAPER},
            "Y" => return &Scissor{name: SCISSOR},
            "Z" => return &Rock{ name: ROCK},
            _ => panic!("Type not convertable")
        }
    }
}

fn convert(str: &str) ->  &'static dyn GameObject{
    match str {
        "A" | "X" => return &Rock{name: ROCK},
        "B" | "Y" => return &Paper{name: PAPER},
        "C" | "Z" => return &Scissor{ name: SCISSOR},
        _ => panic!("Type not convertable")
    }
}

fn part1(lines: Lines) -> i32{
    let mut points : i32 = 0;
    for line in lines {
        let players: Vec<&str> = line.split(' ').collect();
        let player1 = convert(players[0]);
        let player2 = convert(players[1]);
        println!("{0} vs {1}",player1.get_name(), player2.get_name());
        println!("{0} vs {1}",player1.calculate_points(player2), player2.calculate_points(player1));
        points +=  player2.calculate_points(player1);
    }
    return points;
}

fn part2(lines: Lines) -> i32{
    let mut points : i32 = 0;
    for line in lines {
        let players: Vec<&str> = line.split(' ').collect();
        let player1 = convert(players[0]);
        let player2 = player1.convert(players[1]);
        println!("{0} vs {1}",player1.get_name(), player2.get_name());
        println!("{0} vs {1}",player1.calculate_points(player2), player2.calculate_points(player1));
        points +=  player2.calculate_points(player1);
    }
    return points;
}

fn main() {
    println!("Hii");
    let lines: Lines = include_str!("day2_input.txt").lines();
    let lines2: Lines = include_str!("day2_input.txt").lines();
    print!("PART1 TOTAL POINTS: {}", part1(lines));
    print!("PART2 TOTAL POINTS: {}", part2(lines2));
}




#[cfg(test)]
mod tests {
    use crate::{convert, ROCK, PAPER, SCISSOR};

    #[test]
    fn test_converter() {
        assert_eq!(ROCK,convert("A").get_name());
        assert_eq!(ROCK,convert("X").get_name());
        assert_eq!(PAPER,convert("B").get_name());
        assert_eq!(PAPER,convert("Y").get_name());
        assert_eq!(SCISSOR,convert("C").get_name());
        assert_eq!(SCISSOR,convert("Z").get_name());
    }
}
