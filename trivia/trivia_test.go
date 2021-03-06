package trivia

import (
	"testing"

	"github.com/go-chat-bot/bot"
)

func Test_loadScores(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "testload",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loadScores()
		})
	}
}

func Test_saveScores(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "testsave"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loadScores()
			saveScores()
		})
	}
}

func Test_renderScores(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name: "nothing",
		},
	}
	loadScores()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := renderScores()
			if (err != nil) != tt.wantErr {
				t.Errorf("renderScores() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_trivia(t *testing.T) {
	type args struct {
		command *bot.Cmd
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			args: args{
				command: &bot.Cmd{
					User: &bot.User{
						ID:   "fniaodaw",
						Nick: "fesnjfis",
					},
					Command: "answer stuff stuff stuff'",
					Args: []string{
						"answer", "stuff", "stuff", "morestuff",
					},
				},
			},
			want: "Try again...",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := trivia(tt.args.command)
			if (err != nil) != tt.wantErr {
				t.Errorf("trivia() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("trivia() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deepCheckAnswer(t *testing.T) {
	type args struct {
		providedAnswer string
		realAnswer     string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				providedAnswer: "grrgre",
				realAnswer:     "gewge",
			},
			name: "test wrong answer",
			want: false,
		},
		{
			args: args{
				providedAnswer: "THISSHOULDPASS",
				realAnswer:     "thisshouldpass",
			},
			name: "test correct answer",
			want: true,
		},
		{
			args: args{
				providedAnswer: "should",
				realAnswer:     "thisshouldpass",
			},
			name: "test partial match / correct answer",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deepCheckAnswer(tt.args.providedAnswer, tt.args.realAnswer); got != tt.want {
				t.Errorf("deepCheckAnswer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scrubStrings(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				input: `the "<d>\'</d>the (ss) test`,
			},
			name: "scrub test",
			want: "the ss test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scrubStrings(tt.args.input); got != tt.want {
				t.Errorf("scrubStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcAccuracy(t *testing.T) {
	type args struct {
		correct   int
		incorrect int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "float",
			args: args{
				correct:   43,
				incorrect: 88,
			},
			want: "0.328",
		},
		{
			name: "div zero",
			args: args{
				correct:   0,
				incorrect: 0,
			},
			want: "0.000",
		},
		{
			name: "perfect",
			args: args{
				correct:   43,
				incorrect: 0,
			},
			want: "1.000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcAccuracy(tt.args.correct, tt.args.incorrect); got != tt.want {
				t.Errorf("calcAccuracy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_showStats(t *testing.T) {
	type args struct {
		cmd *bot.Cmd
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "show stats",
			args: args{
				cmd: &bot.Cmd{
					User: &bot.User{
						ID: "U6M4BAGHK",
					},
				},
			},
			want: `
Player Name: vikimarras82
Total Score: 94900
Total Correct Answers: 158
Total Wrong Answers: 316
Total New Question Requests: 93
Accuracy: 0.333
`,
		},
	}
	loadScores()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := showStats(tt.args.cmd)
			if (err != nil) != tt.wantErr {
				t.Errorf("showStats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("showStats() = %v, want %v", got, tt.want)
			}
		})
	}
}
