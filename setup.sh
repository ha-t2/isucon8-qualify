cd
git clone https://github.com/tagomoris/xbuild.git
mkdir local
xbuild/go-install 1.10.3 $HOME/local/go
export PATH=$HOME/local/go/bin:$HOME/go/bin:$PATH
(
  cd isucon8-qualify/bench
  make deps
  make
  ./bin/gen-initial-dataset
)
(
  cd isucon8-qualify
  ./db/init-user.sh
  ./db/init.sh
)
(
  cd isucon8-qualify/webapp/perl
  cpanm --installdeps --notest --force .
)


echo "Provisioning Successful"