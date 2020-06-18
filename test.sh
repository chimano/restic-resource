export RESTIC_PASSWORD=testpass
export TEST_REPO="${PWD}/testrepo"
export TEST_OUTFILE="${PWD}/testfile.txt"
export TEST_INDIR="${PWD}/testdir"
export TEST_HOST="testhost"

echo "test" > $TEST_OUTFILE

restic init --repo $TEST_REPO
restic backup --repo testrepo --host $TEST_HOST $TEST_OUTFILE
go test github.com/chimano/restic-resource/cmd/out
go test github.com/chimano/restic-resource/cmd/in
go test github.com/chimano/restic-resource/cmd/check

# cleaning up
rm -rf $TEST_INDIR
rm -rf $TEST_REPO
rm $TEST_OUTFILE

unset RESTIC_PASSWORD
unset TEST_REPO
unset TEST_OUTFILE
unset TEST_INDIR