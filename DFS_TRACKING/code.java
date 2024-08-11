import java.util.Scanner;
import java.awt.*;
import java.awt.event.*;
class Graph extends Frame  
{static int c, n=15, m[][]=new int[n][n], stack[]=new int[n],visited[]=new int[n],v,top=-1,top1=-1;
 static int num=0,stack1[]=new int[n];
    Button b1;
    Label l2;
     Graph()
     {
 		setTitle("dfs");
 		setSize(500,500);
 		setVisible(true);
 		setBackground(Color.BLACK);
     }
     public void paint(Graphics g)
     {   
           g.drawOval(0,450,50,50);
           g.drawOval(200,450,50,50);
           g.drawOval(350,450,50,50);
           g.drawOval(450,450,50,50);
           g.drawOval(200,350,50,50); g.drawOval(200,250,50,50); g.drawOval(200,150,50,50); g.drawOval(200,50,50,50);
                         g.drawOval(290,300,50,50); g.drawOval(290,150,50,50);g.drawOval(350,150,50,50);g.drawOval(450,150,50,50);                                                      
                                                    g.drawOval(290,50,50,50);g.drawOval(350,50,50,50);
                                                                             g.drawOval(350,300,50,50);
            if(num==1)
            { 	
             for(int i=0;i<14;i++)
            { try {Thread.sleep(1000); g.setColor(Color.ORANGE);
            	if(stack1[i]==0)  g.fillOval(0,450,50,50);
            	if(stack1[i]==1)  g.fillOval(200,450,50,50);
            	if(stack1[i]==2)  g.fillOval(350,450,50,50);
            	if(stack1[i]==3)  g.fillOval(450,450,50,50);
            	if(stack1[i]==4)  g.fillOval(200,350,50,50);
            	if(stack1[i]==5)  g.fillOval(200,250,50,50);
            	if(stack1[i]==6)  g.drawOval(290,300,50,50);
            	if(stack1[i]==7)  g.fillOval(200,150,50,50);
            	if(stack1[i]==8) g.drawOval(200,50,50,50);
            	if(stack1[i]==9)  g.drawOval(290,150,50,50);
            	if(stack1[i]==10)   g.drawOval(290,50,50,50);
            	if(stack1[i]==11)  g.drawOval(350,150,50,50);
            	if(stack1[i]==12)  g.drawOval(350,300,50,50);
            	if(stack1[i]==13) ;g.drawOval(350,50,50,50);
            	if(stack1[i]==14) g.drawOval(450,150,50,50);
            }
            catch(Exception k)
            {
            	
            }
            }
            }                                                                
                                                                             
                                                                             
                                                                             
                                                                             
     }

    
     public void pro(int v)
     {   
    	 int count=0;
    	 for(int i=0;i<n;i++)
        {
    	 if((m[v][i]==1)&&(visited[i]==0))
    	 {
    		 top++;top1++;
    		 top=top1;
    		 stack[top]=i;
    		 visited[stack[top]]=1;
    		 count=1;
    		 break;
    	 }
        } 
    	 if(count==0)
    	 {  
    		 top1--;   		
    	 }
    	 if((top1>-1)&&(top<n-1))
    	 {
    		 pro(stack[top1]);
    	 }
    	 
 
      }


	public static void main(String[]args)
	{
		Graph d=new Graph();
		Frame f=new Frame();
		Scanner ab=new Scanner(System.in);
	    System.out.println("enter the matrix");
	    for(int i=0;i<n;i++)
	     for(int j=0;j<n;j++)
	    
	    	m[i][j]=ab.nextInt();
	    	for(int i=0;i<n;i++)
	    	{
	    	    visited[i]=0;
	    	}
	    System.out.println("enter the vetex");
	    v=ab.nextInt();
          top++;top1++;
	     stack[top]=v;
	    visited[stack[top]]=1;
	   	d.pro(v);
 	   	try {
 		   	for(int i=0;i<n;i++)
 		   	{   stack1[i]=stack[i];
 		   	    Thread.sleep(100);
 		   	    num=1;
 		   	    System.out.println(stack[i]);
 		   		if(stack[i]==14)
 		   		{  f.repaint();
 		   			break;
 		   		}
 		   	}
 		   	    }
 		   	catch(Exception ea)
 		   	{
 		   		
 		   	}
	}
}