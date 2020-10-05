package templates

const DefaultPlainVersion = `\documentclass[12pt]{report}
\title{ASSIGNMENT_NAME}
\author{CLASS_NAME \\ DOCUMENT_TYPE \\ CLASS_TEACHER}
\date{DATE}

% ? Packages:
\usepackage{enumitem}  % List Spacing
\usepackage{setspace}  % Set line spacing
\usepackage{url}       % Clickable URLs
\usepackage{color}     % Highlighted text and colored text
\usepackage{soul}      % Highlighted text and colored text
\usepackage{xcolor}    % Highlighted text and colored text
\usepackage{amsmath}   % Advanced mathematics
\usepackage{polynom}   % Advanced mathematics with polynomials
\usepackage{tikz}      % Plotting
\usepackage{pgfplots}  % Plotting
\usepackage{graphicx}  % Loading Images
\usepackage{fancyhdr}  % Page formatting

% ? Configuration:
\setlist[enumerate]{itemsep=0mm}      % Setting enumerate list vertical margins
\setlist[itemize]{itemsep=0mm}        % Setting itemize list vertical margins
\pgfplotsset{compat=1.16}             % Typeset for pgfplots package
\setcounter{secnumdepth}{4}           % Depth of the table of contents
\setcounter{tocdepth}{4}              % Depth of the table of contents
\pagestyle{fancy}                     % Page configuration
\setlength{\headheight}{15pt}         % Page configuration
\fancyhf{}                            % Page configuration
\rfoot{Page \thepage}                 % Page footer configuration

% ? Dynamic Configuration:
\rhead{ASSIGNMENT_NAME $ \cdot $ DOCUMENT_TYPE}
\lhead{AUTHOR_FULL_NAME}
\graphicspath{{../../../../assets/images/PATH_CLASS_NAME/MONTH_NAME/DOCUMENT_TYPE/PATH_ASSIGNMENT_NAME/}}

% ? Commands:
\newcommand{\horizontalrule}{\noindent\makebox[\textwidth]{\rule{\textwidth}{0.6pt}}}   % Horizontal Line
\newcommand{\showimage}[2]{\includegraphics[width=#2]{#1}}                              % Load an image
\newcommand{\tabitem}{~~\llap{\textbullet}~~}                                           % Bullets in tabular environment (tables)
\newcommand{\ctext}[3][RGB]{                                                            % Highlighter base command
    \begingroup
    \definecolor{hlcolor}{#1}{#2}\sethlcolor{hlcolor}
    \hl{#3}
    \endgroup
}

% ? Highlighters:
\newcommand{\green}[1]{\ctext[RGB]{0, 255, 0}{#1}}     % Green
\newcommand{\yellow}[1]{\ctext[RGB]{255, 255, 0}{#1}}  % Yellow
\newcommand{\orange}[1]{\ctext[RGB]{255, 174, 0}{#1}}  % Orange
\newcommand{\red}[1]{\ctext[RGB]{255, 0, 0}{#1}}       % Red

%░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░ START ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
\begin{document}
    \maketitle



\end{document}
%░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░ END ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░

% ? License:
% ! Copyright (c) YEAR_NUMBER AUTHOR_FULL_NAME
`

const DefaultEmojiVersion = `\documentclass[12pt]{report}
\title{ASSIGNMENT_NAME}
\author{CLASS_NAME \\ DOCUMENT_TYPE \\ CLASS_TEACHER}
\date{DATE}

% ? 📦 Packages:
\usepackage{enumitem}  % ↔️ List Spacing
\usepackage{setspace}  % ↔️ Set line spacing
\usepackage{url}       % 🌏 Clickable URLs
\usepackage{color}     % 🟡 Highlighted text and colored text
\usepackage{soul}      % 🟡 Highlighted text and colored text
\usepackage{xcolor}    % 🟡 Highlighted text and colored text
\usepackage{amsmath}   % 🧮 Advanced mathematics
\usepackage{polynom}   % 🧮 Advanced mathematics with polynomials
\usepackage{tikz}      % 📊 Plotting
\usepackage{pgfplots}  % 📊 Plotting
\usepackage{graphicx}  % 📷 Loading Images
\usepackage{fancyhdr}  % 📄 Page formatting

% ? ⚙️ Configuration:
\setlist[enumerate]{itemsep=0mm}      % ↔️ Setting enumerate list vertical margins
\setlist[itemize]{itemsep=0mm}        % ↔️ Setting itemize list vertical margins
\pgfplotsset{compat=1.16}             % 📊 Typeset for pgfplots package
\setcounter{secnumdepth}{4}           % 🧾 Depth of the table of contents
\setcounter{tocdepth}{4}              % 🧾 Depth of the table of contents
\pagestyle{fancy}                     % 📄 Page configuration
\setlength{\headheight}{15pt}         % 📄 Page configuration
\fancyhf{}                            % 📄 Page configuration
\rfoot{Page \thepage}                 % 🏈 Page footer configuration

% ? ⚙️ Dynamic Configuration:
\rhead{ASSIGNMENT_NAME $ \cdot $ DOCUMENT_TYPE}
\lhead{AUTHOR_FULL_NAME}
\graphicspath{{../../../../assets/images/PATH_CLASS_NAME/MONTH_NAME/DOCUMENT_TYPE/PATH_ASSIGNMENT_NAME/}}

% ? 🐚 Commands:
\newcommand{\horizontalrule}{\noindent\makebox[\textwidth]{\rule{\textwidth}{0.6pt}}}   % 📏 Horizontal Line
\newcommand{\showimage}[2]{\includegraphics[width=#2]{#1}}                              % 📷 Load an image
\newcommand{\tabitem}{~~\llap{\textbullet}~~}                                           % 🧾 Bullets in tabular environment (tables)
\newcommand{\ctext}[3][RGB]{                                                            % 🟡 Highlighter base command
    \begingroup
    \definecolor{hlcolor}{#1}{#2}\sethlcolor{hlcolor}
    \hl{#3}
    \endgroup
}

% ? 🟡 Highlighters:
\newcommand{\green}[1]{\ctext[RGB]{0, 255, 0}{#1}}     % 🍏 Green
\newcommand{\yellow}[1]{\ctext[RGB]{255, 255, 0}{#1}}  % 🍋 Yellow
\newcommand{\orange}[1]{\ctext[RGB]{255, 174, 0}{#1}}  % 🍊 Orange
\newcommand{\red}[1]{\ctext[RGB]{255, 0, 0}{#1}}       % 🍒 Red

%░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░🟢 START 🟢░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
\begin{document}
    \maketitle



\end{document}
%░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░🔴 END 🔴░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░

% ? 📜 License:
% ! Copyright (c) YEAR_NUMBER AUTHOR_FULL_NAME
`
