#include <gtest/gtest.h>

TEST(HelloTest, BasicAssertion)
{
    EXPECT_STRNE("hellol", "world");
    EXPECT_EQ(7*6, 42);
}

