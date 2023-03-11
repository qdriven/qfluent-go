
create_issue(){
  TITLE=$1
  FILE_NAME=$2
  CONTENT=`cat docs/${FILE_NAME}`
  echo ${CONTENT}
  echo "\n - [${TITLE}](docs/${FILE_NAME})" >> README.md
  gh issue create --title "${TITLE}" --body "${CONTENT}" --label daily
}

post_issue(){

BASE_PATH='docs/daily/tips/'
FILES=$(ls ${BASE_PATH})
for i in $FILES; do
  TITLE=${i}
  FILE_PATH=${BASE_PATH}${i}
  LABEL="tips"
  echo "cat ${FILE_PATH}"
  content=$(cat ${FILE_PATH})
  echo $content
  gh issue create --title "${TITLE}" --label "${LABEL}" --body "${content}"
done

}