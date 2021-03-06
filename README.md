# HoneyWaffleSDT5

## 振り飛車党評価関数（やねうら王のKPPT形式）
https://drive.google.com/open?id=1mvMtLTy5TM4GFBda-ppX-wEYKylVs7yx

予選リーグ、決勝トーナメントを通して使用した評価関数です。  
定跡で76歩や34歩を突かせると、だいたい飛車を振ります。  
定跡なしの場合、飛車先を突いて居飛車を指してしまうこともよくあります。

## 振り飛車「狂」評価関数（やねうら王のKPPT形式）
https://drive.google.com/open?id=11SSW-2b8VP0YEzEuN1PLwQSNS5FDr7NT

振り飛車党評価関数を作るための原液のような評価関数です。  
ただ飛車を振るだけの評価関数で、とても弱いです。  
やねうら王を魔改造し、居飛車を減点することで強引に振り飛車を高く評価させ、そこから教師局面を生成させて学習しました。  


振り飛車党評価関数は、以下の評価関数を試行錯誤しながらブレンドしたものです。

* 振り飛車「狂」評価関数

* ライブラリ申請した、elmo評価関数

* ライブラリ申請した、やねうら王rezero8評価関数

* やねうら王エンジン+elmo評価関数+振り飛車定跡で教師局面生成し、ゼロベースから学習した評価関数

戦型を気にしなければelmo(WCSC27)に6割以上は勝っていたと思いますが、居飛車で勝たないよう定跡で封印したので、振り飛車のみでは互角くらいかと思います。


## 定跡ファイル（やねうら王形式）

https://github.com/32hiko/HoneyWaffleSDT5/blob/master/waffle_book.db

WCSC27版の定跡をベースにしたものを初期状態としました。  
elmo(WCSC27)との対戦を繰り返し、振り飛車でいい感じの序盤になったら、振り飛車側の手だけ50手まで追加しています。  
また、評価値が急に下がるような局面では局面検討の結果で矯正しています。  
評価値的に多少不利（先手-100まで、後手は-200弱まで）でも、相手から特にいい手がなく差が広がらないものは採用しています。  
開発当初は評価関数も振り飛車党になりきっていないため、定跡を更新しつつ、評価関数も学習しつつのサイクルに乗せるまでは苦労しました。  
級位者がすべて目を通して手作業で追加しているので、不自然なものが結構あるかもしれません。  
評価関数側の成果により、初手58飛や後手での52飛は不要になりました。

## ソースコード

https://github.com/32hiko/HoneyWaffleSDT5/blob/master/evaluate_kppt.cpp

やねうら王での棋譜生成にあたり、魔改造したファイルです。  
当時のやねうら王のソースをこのように改変して棋譜生成することで、振り飛車「狂」を学習させました。  
ドワンゴの電王トーナメント担当者の方に送付した際のコメントをそのまま貼り付けます：

>やねうら王の局面評価のソースです。  
>805-861行目が変更部分です。  
>先手なら2筋、後手なら8筋に飛車がいる状態にペナルティを付与することで、強引に振り飛車を高く評価させます。  
>（ペナルティが28や82のみだと、横歩や相掛かりになり浮き飛車にしてごまかそうとするので、筋で）  
>冗長なロジックですが、学習を回すときにしか動かさないのであまり気にしていません。  
>振り飛車を高く評価したものを学習することで、通常状態のやねうら王でも定跡なしで飛車を振る挙動になります。  

https://github.com/32hiko/HoneyWaffleSDT5/blob/master/waffle_wrapper.go

WCSC27に続き、将棋所にエンジンとして登録する際のラッパーです。  
将棋所 <-> waffle_wrapper <-> やねうら王exe

当初は、定跡の管理をここでやる予定でしたが、評価関数単独でもある程度振り飛車を指すことで、  
定跡にヒットしない場合は58飛車、のような処理をロジックで持たせる必要がなくなったのでやめました。

基本戦略として、飛車を振ることで強制的に対抗形の長い将棋に持ち込めるため、
定跡と合わせ、相手の残り時間を見つつ時間を節約することで、全局を通して時間的優位に立つ作戦です。

ponderのときは怖い、と書いていますが、  
敗勢でのApery戦にて、微妙に持ち時間が数秒残った状態でponderヒットした結果、  
切れ負けとなったためバグがありそうです。

## 例によって自戦記的なまとめ

https://shiroigohanp.tumblr.com/post/167734836519/%E7%AC%AC5%E5%9B%9E%E5%B0%86%E6%A3%8B%E9%9B%BB%E7%8E%8B%E3%83%88%E3%83%BC%E3%83%8A%E3%83%A1%E3%83%B3%E3%83%88honeywaffle%E8%A6%96%E7%82%B9%E3%81%A0%E3%81%91%E3%81%AE%E3%81%BE%E3%81%A8%E3%82%81-%E5%89%8D%E6%97%A5%E6%BA%96%E5%82%99
