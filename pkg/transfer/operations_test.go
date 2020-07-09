package transfer

import (
	"github.com/lozovoya/gohomework5_1/pkg/card"
	"testing"
)

func TestService_Card2Card(t *testing.T) {
	type fields struct {
		CardSvc       *card.Service
		ItoICommision int64
		ItoIMin       int64
		ItoECommision int64
		ItoEMin       int64
		EtoICommision int64
		EtoIMin       int64
		EtoECommision int64
		EtoEMin       int64
	}
	type args struct {
		from   string
		to     string
		amount int64
	}

	CardSvc := card.NewService("qqq")
	CardSvc.IssueCard("master", 100_000_00, "0000 0000 0000 0000", "rub")
	CardSvc.IssueCard("visa", 15_000_00, "3333 3333 3333 3333", "rub")

	tests := []struct {
		name      string
		fields    fields
		args      args
		wantTotal int64
		wantOk    bool
	}{
		{
			name: "ItoI succesful",
			args: args{
				from:   "0000 0000 0000 0000",
				to:     "3333 3333 3333 3333",
				amount: 10_000_00,
			},
			wantTotal: 10_000_00,
			wantOk:    true,
		},
		{
			name: "ItoI failed",
			args: args{
				from:   "0000 0000 0000 0000",
				to:     "3333 3333 3333 3333",
				amount: 10_000_000_00,
			},
			wantTotal: 10_000_000_00,
			wantOk:    false,
		},
		{
			name: "ItoE succesful",
			args: args{
				from:   "0000 0000 0000 0000",
				to:     "3333",
				amount: 10_000_00,
			},
			wantTotal: 10_050_00,
			wantOk:    true,
		},
		{
			name: "ItoE failed",
			args: args{
				from:   "0000 0000 0000 0000",
				to:     "3333",
				amount: 10_000_000_00,
			},
			wantTotal: 10_050_000_00,
			wantOk:    false,
		},
		{
			name: "EtoI",
			args: args{
				from:   "0000",
				to:     "3333 3333 3333 3333",
				amount: 10_000_00,
			},
			wantTotal: 10_000_00,
			wantOk:    true,
		},
		{
			name: "EtoE",
			args: args{
				from:   "0000",
				to:     "3333",
				amount: 10_000_00,
			},
			wantTotal: 10_150_00,
			wantOk:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				CardSvc:       CardSvc,
				ItoICommision: 0,
				ItoIMin:       0,
				ItoECommision: 5,
				ItoEMin:       10_00,
				EtoICommision: 0,
				EtoIMin:       0,
				EtoECommision: 15,
				EtoEMin:       30_00,
			}
			gotTotal, gotOk := s.Card2Card(tt.args.from, tt.args.to, tt.args.amount)
			if gotTotal != tt.wantTotal {
				t.Errorf("Card2Card() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Card2Card() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestService_Transfer(t *testing.T) {
	type fields struct {
		CardSvc       *card.Service
		ItoICommision int64
		ItoIMin       int64
		ItoECommision int64
		ItoEMin       int64
		EtoICommision int64
		EtoIMin       int64
		EtoECommision int64
		EtoEMin       int64
	}
	type args struct {
		from   string
		to     string
		amount int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "ItoI test Error",
			args: args{
				from:   "1111 1111 1111 1111",
				to:     "0000 0000 0000 0000",
				amount: 1_000_000_000_00,
			},
			wantErr: ErrorSourceCardNotEnoughMoney,
		},
	}

	CardSvc := card.NewService("qqq")
	CardSvc.IssueCard("master", 100_000_00, "0000 0000 0000 0000", "rub")
	CardSvc.IssueCard("visa", 15_000_00, "1111 1111 1111 1111", "rub")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				CardSvc:       CardSvc,
				ItoICommision: 0,
				ItoIMin:       0,
				ItoECommision: 5,
				ItoEMin:       10_00,
				EtoICommision: 0,
				EtoIMin:       0,
				EtoECommision: 15,
				EtoEMin:       30_00,
			}
			if err := s.Transfer(tt.args.from, tt.args.to, tt.args.amount); err != tt.wantErr {
				t.Errorf("Transfer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
