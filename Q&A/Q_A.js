let data = [
  {
    title: "这是问题",
    content: "这是答案",
  },
  {
    title: "全国最水的大学是哪所大学？",
    content: "折姜大学",
  },
  {
    title: "这所大学最水的课是哪门课？",
    content: "线代五",
  },
  {
    title: "想不出答案了随便写一个",
    content: "想不出问题了随便写一个",
  },
];
let divs = ["div1", "div2", "div3", "div4"];
let ids = ["q1", "q2", "q3", "q4"];
for (let i = 0; i < 4; ++i)
  document.getElementById(ids[i]).innerHTML = data[i].title;
let showAnswer = [false, false, false, false];
let everClicked = [false, false, false, false];

function Change(num) {
  greyed = everClicked[num];
  expanded = showAnswer[num];
  if (expanded == false) {
    Recover(num);
    let divToShow = document.getElementById(divs[num]);
    let txtToShow = document.getElementById(ids[num]);
    divToShow.style.backgroundColor = "gainsboro";
    txtToShow.style.color = "black";
    divToShow.style.height = "135px";
    txtToShow.innerHTML =
      data[num].title +
      `
                <br>
                <hl> <br>   
            ` +
      data[num].content;
    everClicked[num] = true;
    showAnswer[num] = true;
  } else {
    showAnswer[num] = false;
    let divToHide = document.getElementById(divs[num]);
    let txtToHide = document.getElementById(ids[num]);
    txtToHide.innerHTML = data[num].title;
    divToHide.style.height = "55px";
    divToHide.style.backgroundColor = "#a1a1a1";
  }
}

function Recover(num) {
  for (i = 0; i < 4; ++i) {
    if (i == num) continue;
    if (showAnswer[i] == false) continue;
    divToRecover = document.getElementById(divs[i]);
    txtToRecover = document.getElementById(ids[i]);
    txtToRecover.innerHTML = data[i].title;
    divToRecover.style.height = "55px";
    divToRecover.style.backgroundColor = "#a1a1a1";
  }
}
