package service

import (
	"context"
	"github.com/MenciusCheng/bookmark-manager/bookmark-service/conf"
	"github.com/MenciusCheng/bookmark-manager/bookmark-service/model"
	"github.com/MenciusCheng/bookmark-manager/bookmark-service/util/logging"
	"github.com/MenciusCheng/superman/util/dragons"
	"io/ioutil"
	"os"
	"testing"
)

func Test_strToLink(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{str: `                <DT><A HREF="http://open.163.com/movie/2010/12/G/F/M6UTT5U0I_M6V2T1JGF.html" ADD_DATE="1630981100" ICON="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAYAAAAf8/9hAAAAcElEQVQ4ja1RQQ7AMAhiS///ZXdyMYisy8bJEoqowEccigwgBnHTN6J+zg+KkwYszHet2WSpqC4ywxpMu9gyCCBqArPYd115rNMJ1A64gYzI59s5ZwN3qu9ar0k0pWDcBkrouMf4KeTo4xKnERz/Cy6pAjoL4fNqsQAAAABJRU5ErkJggg==">麻省理工学院公开课：算法导论_课程简介及算法分析_网易公开课</A>`},
			want: true,
		},
		{
			args: args{str: `            <DT><H3 ADD_DATE="1630981781" LAST_MODIFIED="1630981781">软件</H3>`},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strToLink(tt.args.str); got != tt.want {
				t.Errorf("strToLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_strToLinkM(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name      string
		args      args
		wantInfo  model.Link
		wantMatch bool
	}{
		{
			args:      args{str: `                <DT><A HREF="http://open.163.com/movie/2010/12/G/F/M6UTT5U0I_M6V2T1JGF.html" ADD_DATE="1630981100" ICON="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAYAAAAf8/9hAAAAcElEQVQ4ja1RQQ7AMAhiS///ZXdyMYisy8bJEoqowEccigwgBnHTN6J+zg+KkwYszHet2WSpqC4ywxpMu9gyCCBqArPYd115rNMJ1A64gYzI59s5ZwN3qu9ar0k0pWDcBkrouMf4KeTo4xKnERz/Cy6pAjoL4fNqsQAAAABJRU5ErkJggg==">麻省理工学院公开课：算法导论_课程简介及算法分析_网易公开课</A>`},
			wantMatch: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotInfo, gotMatch := strToLinkM(tt.args.str)
			if gotMatch != tt.wantMatch {
				t.Errorf("strToLinkM() gotMatch = %v, want %v", gotMatch, tt.wantMatch)
			}
			t.Logf("%+v", gotInfo)
		})
	}
}

func TestService_InitLinkFromString(t *testing.T) {
	dragons.Init(
		dragons.ConfigPath("../app/config/config.toml"),
	)

	// init local config
	cfg, err := conf.Init()
	if err != nil {
		logging.Fatalf("service config init error %s", err)
	}

	// create a service instance
	s := New(cfg)

	file, err := os.Open("/Users/chengmengwei/Documents/bookmarks_2021_12_7.html")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)

	type args struct {
		ctx context.Context
		str string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				ctx: context.Background(),
				str: string(content),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.InitLinkFromString(tt.args.ctx, tt.args.str); (err != nil) != tt.wantErr {
				t.Errorf("InitLinkFromString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
