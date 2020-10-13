#include <iostream>
#include <algorithm>

using namespace std;

#define MAXWORK 5 //作业数

struct process
{
    int pid;    //作业号
    double come_time;   //到达时
    double run_time;    //运行时
    double begin_time;  //开始时
    double over_time;   //完成时
    double round_time;  //周转时
    double avg_time;    //带权周转时
    double HRR;         //响应比
}pc[MAXWORK];

/**按到时间正序排序*/
bool CmpByComeTime(process p1,process p2){
    return p1.come_time < p2.come_time;
}

/**按id号正序排序*/

bool CmpByPid(process p1,process p2){
    return p1.pid < p2.pid;
}

/**按运行时长正序排序*/

bool CmpByRunTime(process p1,process p2){
    return p1.run_time == p2.run_time ? p1.come_time<p2.come_time : p1.run_time<p2.run_time;
}

/**按响应比逆序排序*/
bool CmpByHRR(process p1,process p2){
    return p1.HRR > p2.HRR;
}

void get_beginAndOver_time(){
    for (int i = 0; i < MAXWORK; i++){
        if(i==0){
            pc[i].begin_time = pc[i].come_time; //第一个作业开始时即为其到达时
        }
        else{
             pc[i].begin_time = pc[i-1].over_time;   
        }
        pc[i].over_time = pc[i].begin_time+pc[i].run_time;
    }
}


void get_roundAndAvg_time() {
	for (int i = 0; i < MAXWORK; ++i)
	{
		pc[i].round_time = pc[i].over_time - pc[i].come_time;		//周转时 = 完成时 - 到达时
		pc[i].avg_time = pc[i].round_time*1.0 / pc[i].run_time;		//平均周转时= 周转时 / 运行时
	}

}

void FCFS(){
    sort(pc,pc+MAXWORK,CmpByComeTime);
    get_beginAndOver_time();
    get_roundAndAvg_time();
}



void print() {
	cout << "执行顺序：" << endl;
	for (int i = 0; i < MAXWORK; ++i)
	{
		/* code */
		cout << pc[i].pid << " ";
	}
	cout << endl;
	cout << "作业号" << '\t' << "到达时" << '\t' << "运行时" << '\t' << "开始时" << '\t' << "完成时" << '\t' << "周转时" << '\t' << "带权周转时" << '\t' << endl;
	sort(pc, pc + MAXWORK, CmpByPid);
	double sum_round_time = 0.0;
	double avg_sum_round_time = 0.0;
	double sum_avg_time = 0.0;
	double avg_sum_avg_time = 0.0;
	for (int i = 0; i < MAXWORK; ++i)
	{
		sum_round_time += pc[i].round_time;
		sum_avg_time += pc[i].avg_time;
 
		cout << pc[i].pid << '\t' << pc[i].come_time << '\t'
			<< pc[i].run_time << '\t' << pc[i].begin_time << '\t'
			<< pc[i].over_time << '\t' << pc[i].round_time << '\t'
			<< pc[i].avg_time << endl;
	}
	avg_sum_round_time = sum_round_time * 1.0 / MAXWORK;
	avg_sum_avg_time = sum_avg_time * 1.0 / MAXWORK;
	cout << "平均周转时: " << avg_sum_round_time << endl
		<< "平均带权周转时: " << avg_sum_avg_time << endl;

}

int main(){
    
    for(int i = 0; i < MAXWORK;i++){
        cin >> pc[i].pid >> pc[i].come_time >> pc[i].run_time;
    }
    FCFS();
    
	cout << "the results of FCFS are:" << endl;
	print();
}