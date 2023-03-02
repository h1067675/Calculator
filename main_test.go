package main

import "testing"

func TestCheckIsNumber(t *testing.T) {
	type args struct {
		a string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Правильный номер",
			args: args{
				a: "1234",
			},
			want: true,
		},
		{
			name: "Ошибочный номер",
			args: args{
				a: "I234",
			},
			want: false,
		},
		{
			name: "Пустое значение",
			args: args{
				a: "",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkIsNumber(tt.args.a)
			if got != tt.want {
				t.Errorf("checkIsNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckIsRoman(t *testing.T) {
	type args struct {
		a string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Просто текст",
			args: args{
				a: "Привет",
			},
			want: false,
		},
		{
			name: "Большое значение с верными символами",
			args: args{
				a: "CMMCXIIL",
			},
			want: true,
		},
		{
			name: "Верное значение",
			args: args{
				a: "XL",
			},
			want: true,
		},
		{
			name: "Пустое значение",
			args: args{
				a: "",
			},
			want: false,
		},
		{
			name: "Арабские цифры",
			args: args{
				a: "1234",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkIsRoman(tt.args.a)
			if got != tt.want {
				t.Errorf("checkIsRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRomanToArabic(t *testing.T) {
	type args struct {
		a string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Верное число 3",
			args: args{
				a: "III",
			},
			want:    3,
			wantErr: false,
		},
		{
			name: "Верное число 499",
			args: args{
				a: "CDXCIX",
			},
			want:    499,
			wantErr: false,
		},
		{
			name: "Верное число 20",
			args: args{
				a: "XX",
			},
			want:    20,
			wantErr: false,
		},
		{
			name: "Верное число 40",
			args: args{
				a: "XL",
			},
			want:    40,
			wantErr: false,
		},
		{
			name: "Верное число 3999",
			args: args{
				a: "MMMCMXCIX",
			},
			want:    3999,
			wantErr: false,
		},
		{
			name: "Ошибочное число 4",
			args: args{
				a: "IIII",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Ошибочное число 4000",
			args: args{
				a: "MMMM",
			},
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := romanToArabic(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("romanToArabic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("romanToArabic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArabicToRoman(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Верное число 1",
			args: args{
				a: 1,
			},
			want: "I",
		},
		{
			name: "Верное число 800",
			args: args{
				a: 800,
			},
			want: "DCCC",
		},
		{
			name: "Верное число 3999",
			args: args{
				a: 3999,
			},
			want: "MMMCMXCIX",
		},
		{
			name: "Верное число 400",
			args: args{
				a: 400,
			},
			want: "CD",
		},
		{
			name: "Верное число 90",
			args: args{
				a: 90,
			},
			want: "XC",
		},
		{
			name: "Верное число 80",
			args: args{
				a: 80,
			},
			want: "LXXX",
		},
		{
			name: "Верное число 17",
			args: args{
				a: 17,
			},
			want: "XVII",
		},
		{
			name: "Верное число 499",
			args: args{
				a: 499,
			},
			want: "CDXCIX",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := arabicToRoman(tt.args.a)
			if got != tt.want {
				t.Errorf("arabicToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculate(t *testing.T) {
	type args struct {
		a int
		b int
		c string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Сложение 1+1",
			args: args{
				a: 1,
				b: 1,
				c: "+",
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "Вычитание 10-1",
			args: args{
				a: 10,
				b: 1,
				c: "-",
			},
			want:    9,
			wantErr: false,
		},
		{
			name: "Умножение 7*3",
			args: args{
				a: 7,
				b: 3,
				c: "*",
			},
			want:    21,
			wantErr: false,
		},
		{
			name: "Деление 8/2",
			args: args{
				a: 8,
				b: 2,
				c: "/",
			},
			want:    4,
			wantErr: false,
		},
		{
			name: "Ошибочный операнд",
			args: args{
				a: 1,
				b: 1,
				c: "=",
			},
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculate(tt.args.a, tt.args.b, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("calculate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
