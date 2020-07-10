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
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr error
	}{
		{
			name: "ItoI test Error",
			args: args{
				from:   "5106 2111 1111 1111",
				to:     "5106 2100 0000 0000",
				amount: 1_000_000_000_00,
			},
			want:    0,
			wantErr: ErrorSourceCardNotEnoughMoney,
		},
		{
			name: "ItoI, from card not found",
			args: args{
				from:   "5106 2100 1111 1111",
				to:     "5106 2100 0000 0000",
				amount: 1_000_000_00,
			},
			want:    0,
			wantErr: ErrorSourceCardNotFound,
		},
		{
			name: "ItoI, to card not found",
			args: args{
				from:   "5106 2111 1111 1111",
				to:     "5106 2111 0000 0000",
				amount: 1_000_000_00,
			},
			want:    0,
			wantErr: ErrorDestCardNotFound,
		},
		{
			name: "ItoI, succesful",
			args: args{
				from:   "5106 2111 1111 1111",
				to:     "5106 2100 0000 0000",
				amount: 1_000_00,
			},
			want:    1_000_00,
			wantErr: nil,
		},
	}

	CardSvc := card.NewService("qqq")
	CardSvc.IssueCard("master", 100_000_00, "5106 2100 0000 0000", "rub")
	CardSvc.IssueCard("visa", 15_000_00, "5106 2111 1111 1111", "rub")

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
			got, err := s.Card2Card(tt.args.from, tt.args.to, tt.args.amount)
			if err != tt.wantErr {
				t.Errorf("Card2Card() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Card2Card() got = %v, want %v", got, tt.want)
			}
		})
	}
}
