export RESTIC_PASSWORD=testpass
export TEST_REPO="${PWD}/testrepo"
export TEST_FILE="${PWD}/testfile.txt"

echo "test" > $TEST_FILE

restic init --repo $TEST_REPO
go test github.com/chimano/restic-resource/cmd/out
rm -rf $TEST_REPO
rm $TEST_FILE