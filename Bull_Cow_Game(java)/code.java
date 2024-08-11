import javax.swing.*;
import java.awt.*;
import java.awt.event.*;
import java.util.Random;
public class COW {
    public static void main(String[] args) {
        JFrame frame = new JFrame("Main Frame");
        frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
      
        JButton button = new JButton("click here for instucrions");
        button.addActionListener(new ActionListener() {
            
            public void actionPerformed(ActionEvent e) {
                JFrame newFrame = new JFrame("New Frame");
                newFrame.setDefaultCloseOperation(JFrame.DISPOSE_ON_CLOSE);
                newFrame.setSize(1000, 1000);
                newFrame.setVisible(true);
                
                JPanel panel = new JPanel() 
                        {
                     protected void paintComponent(Graphics g) {
                         super.paintComponent(g);
                         Image image = new ImageIcon("C:/Users/hi/Downloads/cowandbull2.jpg").getImage();
                         g.drawImage(image, 0, 0, getWidth(), getHeight(), this);
                     
                     
                        }
                        };
                    newFrame.add(panel);
                    JButton button1 = new JButton("click here to play the game");
                     button1.addActionListener(new ActionListener() {
                         
                         public void actionPerformed(ActionEvent e) {
                             JFrame newFrame = new JFrame("New Frame");
                             newFrame.setDefaultCloseOperation(JFrame.DISPOSE_ON_CLOSE);
                             newFrame.setSize(1000, 1000);
                             newFrame.setVisible(true);
                             newFrame.setLayout(null);
                             //JLabel background = new JLabel(new ImageIcon("C:/Users/hi/Downloads/cowbull.jpg"));
                          
                             //background.setBounds(0, 0, newFrame.getWidth(), newFrame.getHeight());

                             
                             //newFrame.getContentPane().add(background);

                             
                             //newFrame.getContentPane().setComponentZOrder(background, 0);
                             final JLabel l,l2,l3;
                             final JButton b;
                             final JTextField tf;
                         
                             String[] possibleWords = {"lady", "dog", "cat", "bird", "lion", "tiger", "horse", "snake"};
                             Random random = new Random();
                             final String word = possibleWords[random.nextInt(possibleWords.length)];
                              
                             b=new JButton("Enter");
                             tf=new JTextField(4);
                             l=new JLabel("enter the word");
                             
                             l.setFont(new Font("Arial", Font.BOLD, 18));
                             l.setForeground(Color.black);
                             
                           
                             l2=new JLabel();
                             
                             l2.setFont(new Font("Arial", Font.BOLD, 18));
                             l2.setForeground(Color.black);
                             
                           
                             
                             l.setBounds(100,250,300,100);
                             tf.setBounds(250,200,100,50);
                             b.setBounds( 200,300,70,70);
                             l2.setBounds(100,350,500,500); 
                             newFrame.add(b);
                             newFrame.add(tf);newFrame.add(l);newFrame.add(l2);
                             JPanel panel = new JPanel() 
                             {
                          protected void paintComponent(Graphics g) {
                              super.paintComponent(g);
                              Image image = new ImageIcon("C:\Users\Lenovo\OneDrive\Desktop\Java Project\mid2\OOPs").getImage();
                              g.drawImage(image, 0, 0, getWidth(), getHeight(), this);
                          
                          
                             }
                             };
                         newFrame.add(panel);
                             b.addActionListener(new ActionListener() {
                                 int count=15,k=0;
                                 public void actionPerformed(ActionEvent e)
                                 {
                                     
                                   int cow=0,i,j;
                                   int bull=0;
                                   String s1=tf.getText();
                                   if(s1.length()<4||s1.length()>4)
                                       JOptionPane.showMessageDialog(null, "please enter a four letter word :-)");
                                       else if((k==0)&&(count>0))
                                       {
                                           for(i=0;i<4;i++)
                                           {
                                               for(j=0;j<4;j++)
                                               {
                                                   if(word.charAt(i)==s1.charAt(j))
                                                   {
                                                       if(i==j)
                                                       cow++;
                                                       else
                                                       bull++;
                                                   }
                                               }
                                           }
                                           
                                           if(cow==4)
                                           {
                                               JOptionPane.showMessageDialog(null, "Congratulations! You won the game! with score"+count);
                                               k=1;
                                           } 
                                           count--;
                                           l2.setText("cows="+cow+ "bulls="+bull+"number of tries left"
                                                + "="+count);
                                            if(count<1)
                                           {
                                                JOptionPane.showMessageDialog(null, "Sorry, You Lost the Game. Better Luck Next Time!");
                                              
                                           }
                                           
                                       } 
                             
                                 }
                                    
                                    
                             });
                             panel.add(this);
                             
                         }});
                    panel.add(button1);
                
            }
        });
        
        JPanel panel = new JPanel() 
                {
             protected void paintComponent(Graphics g) {
                 super.paintComponent(g);
                 Image image = new ImageIcon("C:/Users/hi/Downloads/cowandbull.jpeg").getImage();
                 g.drawImage(image, 0, 0, getWidth(), getHeight(), this);
             
             
                }};
            panel.add(button);
        frame.add(panel);
        frame.setSize(1000, 1000);
        frame.setVisible(true);
        
    

        }
}