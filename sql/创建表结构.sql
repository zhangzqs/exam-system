CREATE TABLE users
(
    uid      SERIAL PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    password TEXT        NOT NULL
);
CREATE TYPE QuestionType as ENUM('single','multiple','fill','judge');

CREATE TABLE questions
(
    qid        SERIAL PRIMARY KEY,
    created_by INT REFERENCES users (uid) NOT NULL,
    title      TEXT                       NOT NULL,
    type       QuestionType               NOT NULL,
    -- 如果是选择题，则不为空，其他题型留空
    -- 选择题的存储结构以json字符串的方式存放
    option     TEXT,
    -- 可能有的题目就没答案
    -- 若为选择题，则为json数字数组字符串
    -- 若为填空题，则为json字符串数组字符串
    -- 若为判断题，则为文本T or F
    answer     TEXT
);
CREATE TABLE papers
(
    pid        SERIAL PRIMARY KEY,
    created_by INT REFERENCES users (uid) NOT NULL,
    title      TEXT                       NOT NULL
);
CREATE TABLE paper_question
(
    qid   int REFERENCES questions (qid),
    pid   int REFERENCES papers (pid),
    score int
);
CREATE TABLE room
(
    -- 一次测验
    rid        SERIAL PRIMARY KEY,
    -- 试卷号
    pid        INT REFERENCES papers (pid) NOT NULL,
    -- 预计开始与截止时间
    start_time TIMESTAMP WITH TIME ZONE    NOT NULL,
    end_time   TIMESTAMP WITH TIME ZONE    NOT NULL
);
CREATE TABLE user_room
(
    uid       INT REFERENCES users (uid) NOT NULL,
    rid       INT REFERENCES room (rid)  NOT NULL,

    -- 进入和提交的时间，可空
    enter_at  TIMESTAMP WITH TIME ZONE,
    submit_at TIMESTAMP WITH TIME ZONE,

    -- 教师评语及分数，可空
    comment   TEXT,
    score     REAL
);
CREATE TABLE user_answer
(
    uid         INT REFERENCES users (uid)     NOT NULL,
    -- 测验号
    rid         INT REFERENCES room (rid)      NOT NULL,
    -- 问题编号
    qid         INT REFERENCES questions (qid) NOT NULL,
    -- 用户的答案
    -- 若为NULL则表示没写
    user_answer TEXT
);
