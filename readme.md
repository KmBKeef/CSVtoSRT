# A tool to convert CSV files into SRT


CSV file structure (input) should look like the following:  


| ID | start time   | end time    | "Subtitle"                          |
|----|--------------|-------------|-------------------------------------|
| 1  | 00:00:00,000 | 00:00:30,00 | "This is a subtitle blah blah blah" |
| 2  | 00:00:30,01  | 00:01:00,00 | "The following part blah blah blah" |


Output will look like:

```sh
1
00:00:00,00 --> 00:00:30,00
this is a subtitle blah blah blah


2
00:00:30,01 --> 00:01:00,00
 The following part blah blah blah

```

## TODO
- [X] Read filename from the user.
- [ ] cut CSV data in blocks
- [ ] Output them in SRT format
