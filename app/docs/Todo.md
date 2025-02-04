## 実装の順番

- domain
- usecase
- infra
- config
- presen
- server
- cmd
- docs

## todo

- sqlcのgroupにeventsを追加したり、findAllでgroupIDとかもとれるように修正する
- resolver->usecase->repoの順でデータモデルを修正していく
- resolver,usecase,repoでsave,updateのときにuserやgroup,eventを返すようにする
- middlewareのコメントの部分を実装
- configがコピペしただけのやつだからコードに合うように変更
- unit testをする